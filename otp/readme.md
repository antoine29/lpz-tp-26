# otp be
on this folder you need folling files to be mounted as a volume on the OTP container:

- a gtfs.zip file
- a city .pbf file
- the entrypoint.sh file (used to hijack the default container entrypoint)

docker run \
-v /home/anthony/lpz-tp/otp/:/var/otp/v3-prod/waltti \
-p 8081:8080 \
-d \
-e OTP_GRAPH_DIR=v3-prod/waltti \
--entrypoint /var/otp/v3-prod/waltti/entrypoint.sh \
hsldevcom/opentripplanner:v2-prod-waltti

docker run \
-v /home/anthony/prjs/lpz-tp-26/otp/:/var/otp/v3-prod/waltti \
-p 8081:8080 \
-d \
-e OTP_GRAPH_DIR=v3-prod/waltti \
--entrypoint /var/otp/v3-prod/waltti/entrypoint.sh \
hsldevcom/opentripplanner:v2-prod-waltti

