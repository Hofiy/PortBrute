package plugins_test

import (
	"asset-scan/plugins"
	"testing"
)

func TestScanMssql(t *testing.T) {
	t.Log(plugins.ScanMssql("127.0.0.1", "1433", "sa", "123456"))
}
