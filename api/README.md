

**to get started**

1. install docker
2. ./build.sh
3. cd init && do run init.go
4. ./run.sh


**Example**

*To query for a token*
```bash
curl "http://localhost:8180/open/login?api=test&email=user1@example.com&password=supertest"
```

result
```json
{
"result":
"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcGkiOiJ0ZXN0IiwiZXhwIjoxNDE1NTEyMTQ2LCJ1c2VyIjoiYzA4NjAzZjgtYTFiZC00OTI1LTg5ZjQtZGY5MWI5MWMyZThlIn0.Ypa7SFqtFzcKxRDwjA8-qFEp2gX1nE8wfnVdtU422R1ykXUQiwnk2Er0piEMiGZORQEdcQrnr1GsRgYGc9PSzOIwsMpltZC_ikhMBqD9AjWNnBp_g_CFGgYW9DFwKbLVDsWKCOZUvXMYKmKp7w2v0NdoeLgLiBQqrsi4TOXbH2TB9rykV1mY2S8x5Qw4nK-niNOrODYB2U2MfKH9UU5zCwAMLM-CCi012xJnMOwXDqlMS3UoEpfBX2xCsMuTAynnurD174UiWlY6v4bnyoK2ZU5i-O0q2is8OGt6IKRSOVz7tqfu1ogBujdV6upGSFoZREF-WQD6U9mF_6_X1HM6eQ"
}
```

*to get the profile of the logged in user*
```bash
curl -XGET \
"http://localhost:8180/user/profile" \
-H "Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcGkiOiJ0ZXN0IiwiZXhwIjoxNDE1NTEyMTQ2LCJ1c2VyIjoiYzA4NjAzZjgtYTFiZC00OTI1LTg5ZjQtZGY5MWI5MWMyZThlIn0.Ypa7SFqtFzcKxRDwjA8-qFEp2gX1nE8wfnVdtU422R1ykXUQiwnk2Er0piEMiGZORQEdcQrnr1GsRgYGc9PSzOIwsMpltZC_ikhMBqD9AjWNnBp_g_CFGgYW9DFwKbLVDsWKCOZUvXMYKmKp7w2v0NdoeLgLiBQqrsi4TOXbH2TB9rykV1mY2S8x5Qw4nK-niNOrODYB2U2MfKH9UU5zCwAMLM-CCi012xJnMOwXDqlMS3UoEpfBX2xCsMuTAynnurD174UiWlY6v4bnyoK2ZU5i-O0q2is8OGt6IKRSOVz7tqfu1ogBujdV6upGSFoZREF-WQD6U9mF_6_X1HM6eQ"
```
result
```json
{"result":{
"id":"c08603f8-a1bd-4925-89f4-df91b91c2e8e",
"dob":"0001-01-01T00:00:00Z",
"email":"user1@example.com",
"created":"0001-01-01T00:00:00Z",
"updated":"0001-01-01T00:00:00Z"
}}
```
