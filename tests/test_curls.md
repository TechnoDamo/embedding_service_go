curl -X POST "http://localhost:5001/v1/chunk/hybrid/file" \
  -H "Content-Type: multipart/form-data" \
  -H "Authorization: APIKeyAuth" \
  -F "files=@Test.pdf" 

curl -X POST "http://localhost:5001/v1/convert/file" \
  -H "Content-Type: multipart/form-data" \
  -H "Authorization: APIKeyAuth" \
  -F "files=@Test.pdf" \
  -F "convert_do_ocr=false"

curl -X POST "http://localhost:5001/v1/convert/file/async" \
  -H "Content-Type: multipart/form-data" \
  -H "Authorization: APIKeyAuth" \
  -F "files=@Test.pdf" \
  -F "convert_do_ocr=false"

curl -X GET "http://localhost:5001/v1/status/poll/{730a4684-60cd-4861-a735-83d19c837ef9}" 

curl -X GET "http://localhost:5001/v1/result/{730a4684-60cd-4861-a735-83d19c837ef9}" 