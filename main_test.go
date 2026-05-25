package main

import "testing"

func TestSimpleFactory(t *testing.T) {
    f := SimpleFactory("http://localhost")
    
    if f.Url != "http://localhost" {
        t.Errorf("feature incorrect, got %s, want: %s", f.Url, "http://localhost")
    }
}

func TestSimpleFactoryEmpty(t *testing.T) {
    f := SimpleFactory("")
    
    if f.Url != "" {
        t.Errorf("expected empty URL for empty host, got %s", f.Url)
    }
}