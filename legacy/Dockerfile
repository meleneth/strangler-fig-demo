FROM python:3.14.4-trixie
WORKDIR /app
COPY . .
RUN pip install --no-cache-dir . -r demo/requirements.txt
EXPOSE 8000
CMD ["legacy"]

LABEL org.opencontainers.image.vendor="Mobilis"
LABEL mobilis.project="http-migration"
LABEL mobilis.service="legacy"
