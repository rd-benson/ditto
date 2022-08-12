package main

import "testing"

// Easylog message output examples
// With publish seperate set
// {
// 		"date": "2022-08-10T07:56:55.0Z",
// 		"id": "AAAAAA010000000002100466",
// 		"device" : "lighting",
// 		"vars": [
// 			{ "name": "power", "value": "1.23"},
// 			{ "name": "current", "value": "4.56"}
// 			]
// 	}

// Without publish separate set
// {
// 		"date": "2022-08-10T07:56:55.0Z",
// 		"id": "AAAAAA010000000002100466",
// 		"vars": [
// 			{ "name": "lighting < power", "value": "1.23"},
// 			{ "name": "lighting < current", "value": "4.56"}
// 			]
// 	}

func TestMqttToPoint(t *testing.T) {

}
