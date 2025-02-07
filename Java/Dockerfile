FROM openjdk:17-slim

COPY MCM.java /app/MCM.java

WORKDIR /app

RUN javac MCM.java
CMD ["java", "MCM"]
