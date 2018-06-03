package tcloud

import (
	"testing"
	"io/ioutil"
	"encoding/base64"
	"os"
)

func TestClient_CompositeImage(t *testing.T) {
	client := NewClient(11, "", 3, 9, 100, 60)

	data, err := ioutil.ReadFile("../test_materials/face.jpg")
	if err != nil {
		t.Fatal(err)
	}

	req := &FaceFusionRequest{4, []byte(base64.StdEncoding.EncodeToString(data))}

	res, err := client.FaceFusion(req)
	if err != nil {
		t.Fatal(err)
	}

	if res.Code != 0 {
		t.Fatalf("code: %d, msg: %s", res.Code, res.Msg)
	}

	d, err := base64.StdEncoding.DecodeString(res.Data.Image)
	if err != nil {
		t.Fatal(err)
	}
	ioutil.WriteFile("../test_materials/face_output.jpg", d, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
}
