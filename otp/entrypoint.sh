#!/bin/bash
set -e

MODE="${1:-}"

echo "OTP_GRAPH_DIR: ${OTP_GRAPH_DIR}"
echo "MODE: ${MODE:-build-serve}"

case "${MODE}" in
  build)
    echo "Building OTP graph..."
    java $JAVA_OPTS -cp @/app/jib-classpath-file @/app/jib-main-class-file /var/otp/${OTP_GRAPH_DIR} --build --save
    ;;
  server)
    echo "Starting OTP server..."
    java $JAVA_OPTS -cp @/app/jib-classpath-file @/app/jib-main-class-file /var/otp/${OTP_GRAPH_DIR} --load --serve
    ;;
  *)
    echo "Building and serving OTP..."
    java $JAVA_OPTS -cp @/app/jib-classpath-file @/app/jib-main-class-file /var/otp/${OTP_GRAPH_DIR} --build --serve
    ;;
esac

