package builtins

import (
    "os"
    "path/filepath"  
    "testing"
)

func TestMakeDirectory(t *testing.T) {
    // Temporary directory to contain the test directory structure
    tmpDir := t.TempDir()

    // Define test cases
    tests := []struct {
        name    string
        dirPath string
        mode    os.FileMode
        wantErr bool
    }{
        {"CreateDir", "testdir", 0755, false},
        {"CreateNestedDir", "nested/dir", 0755, false},
        {"CreateDirNoPermission", "/invalid/dir", 0755, true},
        // More test cases...
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            fullPath := filepath.Join(tmpDir, tt.dirPath)
            err := MakeDirectory(fullPath, tt.mode)  // Remove 'builtins.'
            if (err != nil) != tt.wantErr {
                t.Errorf("MakeDirectory() error = %v, wantErr %v", err, tt.wantErr)
            }

            if !tt.wantErr {
                if _, err := os.Stat(fullPath); os.IsNotExist(err) {
                    t.Errorf("MakeDirectory() did not create directory")
                }
            }
        })
    }
}
