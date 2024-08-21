package generate

import (
	"testing"
)

func TestConvertJsonSchema2Struct(t *testing.T) {
	schema := `
{
   "$schema":"https://json-schema.org/draft/2020-12/schema",
   "$id":"https://fuxi.byted.org/schema/video_object.json",
   "title":"TikTok Video Object NativeSchema",
   "description":"...",
   "type":"object",
   "properties":{
      "OriginalVideo":{
         "type":"object",
         "properties":{
            "Manifest":{
               "$ref":"#/definitions/manifest"
            }
         }
      },
      "EncodedStreams":{
         "type":"object",
         "additionalProperties": {"$ref":"#/definitions/manifest"}
      },
      "Size":{
         "type":"integer"
      },
      "IsEncrypted":{
         "type":"boolean"
      },
      "EncryptionEKey":{
         "type":"string"
      }
   },
   "definitions":{
      "manifest":{
         "type":"object",
         "Properties":{
            "Handle":{
               "type":"object",
               "additionalProperties": {"type":"string"}
            },
            "Reference":{
               "type":"object",
               "additionalProperties": {"type":"string"}
            }
         }
      }
   }
}
`
	result, err := ConvertJsonSchema2Struct(schema, "model")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("result: \n%s", result)
}
