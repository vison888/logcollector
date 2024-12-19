@echo on
protoc --go_out=./log_collector --vkit_out=./log_collector/ --vkit_opt=--handlePath=../handler --validate_out="lang=go:./log_collector"  .\log_collector\log_collector.proto
exit

