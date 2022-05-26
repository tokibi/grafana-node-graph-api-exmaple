# grafana-node-graph-api-exmaple

## Running Grafana and API Server

```
docker-compose up -d
```

## Setup data source

1. Open `http://localhost:3000` in your browser.
2. Login grafana dashboard (admin/admin)   
3. Move to `Configuration -> Data sources` and Click `Add data source`
4. Select `Node Graph API`
   1. Enter `http://data-api:8080` in URL field.
   2. Click `Save & Test`
