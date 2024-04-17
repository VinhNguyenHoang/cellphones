FROM bitnami/kafka:3.7.0-debian-12-r2

RUN mkdir -p /bitnami/kafka/plugins

COPY /deployments/mongo-kafka-connect-1.11.2-all.jar ./opt/bitnami/kafka/libs

# COPY /deployments/debezium-connector-mongodb /deployments/debezium-connector-jdbc ./opt/bitnami/kafka/libs/

CMD /opt/bitnami/kafka/bin/connect-standalone.sh /bitnami/kafka/config/connect-standalone.properties