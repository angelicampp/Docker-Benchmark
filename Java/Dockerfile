FROM openjdk:17-slim

COPY MCM.java /MCM.java

RUN javac MCM.java
CMD ["sh", "-c", "java MCM > /app/output.txt"]
