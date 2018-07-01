/home/espadon/programming/golang/gocode/bin/golint -set_exit_status
retVal=$?
if [ $retVal -eq 1 ]; then
  echo "linter check is failed"
  exit $retVal
fi
go build main.go
go test ./unit_tests -run ''
unitRetVal=$?
if [ $unitRetVal -ne 0 ]; then
  echo "unit tests fail"
  exit $unitRetVal
fi
#cd /home/espadon/vgnt
#vagrant reload test_api --provision