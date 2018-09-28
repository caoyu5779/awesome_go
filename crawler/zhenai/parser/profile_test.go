package parser

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"selfLearning/crawler/model"
	"testing"
)

func TestParseProfile(t *testing.T) {
	//contents, err := fetcher.Fetch("http://album.zhenai.com/u/1008140144")
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "安静的雪")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1"+"Element; but was : %v", result.Items)
	}

	profile := result.Items[0]
	fmt.Println(profile)
	expected := model.Profile{
		Age:        41,
		Height:     1,
		Weight:     1,
		Income:     "",
		Gender:     "",
		Name:       "",
		Xinzuo:     "",
		Occupation: "",
		Marriage:   "",
		House:      "",
		Hokou:      "",
		Education:  "",
		Car:        "",
	}

	if !reflect.DeepEqual(profile, expected) {
		t.Errorf("got : %v, want : %v", profile, expected)
	}
}
