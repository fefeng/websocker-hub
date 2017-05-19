FROM alpine
RUN mkdir /app 
RUN mkdir /app/ui 
WORKDIR /app
COPY websocker-hub /app
COPY ui /app/ui
# ENV MODULE_VERSION #MODULE_VERSION#

ENTRYPOINT ["/app/websocker-hub"]
CMD ["--port=18886"]
