import express from "express";
const app = express();
const clientsRouter = require("./api/clients");

const usersRouter = require("./api/users");
const nutritionistsRouter = require("./api/nutritionists");

app.use(express.json());

app.use("/clients", clientsRouter);
app.use("/nutritionists", nutritionistsRouter);
app.use("/users", usersRouter);

app.listen(3000, () => console.log("Holiwi 🥑"));
