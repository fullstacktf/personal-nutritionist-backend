import express from "express";
const app = express();
const clientsRouter = require("./api/clients");

app.use(express.json());
app.use("/clients", clientsRouter);

app.listen(3000);
