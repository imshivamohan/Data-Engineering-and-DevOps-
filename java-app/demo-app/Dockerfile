# ===== Builder Stage =====
FROM maven:3.8.7-amazoncorretto-17 AS builder

WORKDIR /app
COPY . .

RUN mvn clean package -DskipTests

# ===== JRE Minimal Stage =====
FROM amazoncorretto:17-alpine AS corretto-jdk

# Required for strip-debug to work
RUN apk add --no-cache binutils

# Build a minimal JRE image
RUN jlink \
         --verbose \
         --add-modules java.base,java.compiler,java.desktop,java.instrument,java.management,java.naming,java.prefs,java.rmi,java.scripting,java.security.jgss,java.sql,jdk.httpserver,jdk.jfr,jdk.unsupported,jdk.crypto.ec,jdk.crypto.cryptoki \
         --strip-debug \
         --no-man-pages \
         --no-header-files \
         --compress=2 \
         --output /jre

# ===== Final Runtime Stage =====
FROM alpine:latest

ENV JAVA_HOME=/jre
ENV PATH="${JAVA_HOME}/bin:${PATH}"

COPY --from=corretto-jdk /jre $JAVA_HOME

EXPOSE 8080

# Use wildcard to copy any JAR file that exists
COPY --from=builder /app/target/*.jar /app/greetings.jar

WORKDIR /app

CMD ["java", "-jar", "greetings.jar"]

