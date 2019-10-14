FROM golang:1.11-alpine

COPY demomesher /opt/demo-mesher/bin/
COPY start.sh /opt/demo-mesher/bin/


WORKDIR /opt/demo-mesher/bin/
RUN chmod +x /opt/demo-mesher/bin/start.sh
CMD ["sh", "-c", "/opt/demo-mesher/bin/start.sh"]