# otp be

docker run \
-v /home/anthony/lpz-tp/otp/:/var/otp/v3-prod/waltti \
-p 8081:8080 \
-d \
-e OTP_GRAPH_DIR=v3-prod/waltti \
--entrypoint /var/otp/v3-prod/waltti/entrypoint.sh \
hsldevcom/opentripplanner:v2-prod-waltti

