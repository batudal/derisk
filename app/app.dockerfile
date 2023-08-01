FROM alpine:latest
COPY de-risk .
COPY views /views
COPY public /public
CMD [ "./de-risk"]
