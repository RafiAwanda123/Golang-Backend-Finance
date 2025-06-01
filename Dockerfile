FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/finance-manager .
COPY --from=builder /app/.env .  
EXPOSE 8080
CMD ["./finance-manager"]