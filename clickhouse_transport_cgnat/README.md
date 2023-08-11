# Description

The transport is inspired by [this
fork](https://github.com/h4sh5/goflow-clickhouse), however it doesn't account
for the routers that send samples frequently in smaller chunks. MergeTree()
performs badly on multiple INSERTS. This attempt is to buffer the incoming data
and perform bulk writes to Clickhouse.
