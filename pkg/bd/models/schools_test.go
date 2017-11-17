package models


import (
  "testing"
)

func TestGetTableNameSchools(t *testing.T) {
  out := Schools{}
  expectedValue := "schools"
  actualValue := out.GetTableName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFileNameSchools(t *testing.T) {
  out := Schools{}
  expectedValue := "Schools.csv"
  actualValue := out.GetFileName()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}

func TestGetFilePathSchools(t *testing.T) {
  out := Schools{}
  expectedValue := "/Users/robertrowe/src/baseballdatabank/core/Schools.csv"
  actualValue := out.GetFilePath()

  if actualValue != expectedValue {
    t.Errorf("actualValue (%s) != expectedValue (%s)\n", actualValue, expectedValue)
  }
}
