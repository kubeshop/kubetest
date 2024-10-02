package artifacts

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"path/filepath"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/kubeshop/testkube/pkg/cloud/data/testworkflow"

	"github.com/kubeshop/testkube/cmd/testworkflow-toolkit/env"
	cloudexecutor "github.com/kubeshop/testkube/pkg/cloud/data/executor"
	"github.com/kubeshop/testkube/pkg/filesystem"
	"github.com/kubeshop/testkube/pkg/ui"
)

// JUnitPostProcessor is a post-processor that checks XML files for JUnit reports and sends them to the cloud.
type JUnitPostProcessor struct {
	fs         filesystem.FileSystem
	client     cloudexecutor.Executor
	root       string
	pathPrefix string
}

func NewJUnitPostProcessor(fs filesystem.FileSystem, client cloudexecutor.Executor, root, pathPrefix string) *JUnitPostProcessor {
	return &JUnitPostProcessor{fs: fs, client: client, root: root, pathPrefix: pathPrefix}
}

func (p *JUnitPostProcessor) Start() error {
	return nil
}

// Add checks if the file is a JUnit report and sends it to the cloud.
func (p *JUnitPostProcessor) Add(path string) error {
	uploadPath := path
	if p.pathPrefix != "" {
		uploadPath = filepath.Join(p.pathPrefix, uploadPath)
	}
	absPath := path
	if !filepath.IsAbs(path) {
		absPath = filepath.Join(p.root, absPath)
	}
	file, err := p.fs.OpenFileRO(absPath)
	if err != nil {
		return errors.Wrapf(err, "failed to open %s", path)
	}
	defer func() { _ = file.Close() }()
	stat, err := file.Stat()
	if err != nil {
		return errors.Wrapf(err, "failed to get file info for %s", path)
	}
	if ok := isXMLFile(stat); !ok {
		return nil
	}
	xmlData, err := io.ReadAll(file)
	if err != nil {
		return errors.Wrapf(err, "failed to read %s", path)
	}
	fmt.Printf("================ XML Data Start ================\n%s\n================ XML Data End ================\n", xmlData)
	ok := isJUnitReport(xmlData)
	if !ok {
		fmt.Printf("Skipping non-JUnit report: %s\n", ui.LightCyan(path))
		return nil
	}
	fmt.Printf("Processing JUnit report: %s\n", ui.LightCyan(path))
	if err := p.sendJUnitReport(uploadPath, xmlData); err != nil {
		return errors.Wrapf(err, "failed to send JUnit report %s", stat.Name())
	}
	return nil
}

// sendJUnitReport sends the JUnit report to the Agent gRPC API.
func (p *JUnitPostProcessor) sendJUnitReport(path string, report []byte) error {
	// Apply path prefix correctly
	if p.pathPrefix != "" {
		path = filepath.Join(p.pathPrefix, path)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	fmt.Printf("Sending JUnit report to the cloud: %s\n", ui.LightCyan(path))
	_, err := p.client.Execute(ctx, testworkflow.CmdTestWorkflowExecutionAddReport, &testworkflow.ExecutionsAddReportRequest{
		ID:           env.ExecutionId(),
		WorkflowName: env.WorkflowName(),
		WorkflowStep: env.Ref(), // TODO: think if it's valid for the parallel steps that have independent refs
		Filepath:     path,
		Report:       report,
	})
	return err
}

// isXMLFile checks if the file is an XML file based on the extension.
func isXMLFile(stat fs.FileInfo) bool {
	if stat.IsDir() || stat.Size() == 0 {
		return false
	}

	return strings.HasSuffix(stat.Name(), ".xml")
}

// isJUnitReport checks if the XML data is a JUnit report.
func isJUnitReport(xmlData []byte) bool {
	const BYTE_SIZE_8KB = 8 * 1024

	tags := []string{
		"<testsuite",
		"<testsuites",
	}

	if len(xmlData) > BYTE_SIZE_8KB {
		xmlData = xmlData[:BYTE_SIZE_8KB]
	}

	content := string(xmlData)

	for _, tag := range tags {
		if !strings.Contains(content, tag) {
			return false
		}
	}

	return true
}

func (p *JUnitPostProcessor) End() error {
	return nil
}
