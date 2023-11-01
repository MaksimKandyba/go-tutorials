POSTFIX=`go list -f '{{.Target}}'`
POSTFIX=${POSTFIX%/*}
export PATH=$PATH:$POSTFIX

