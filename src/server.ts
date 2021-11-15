import express from "express";
const app = express();

const usersRouter = require("./api/users");
const nutritionistsRouter = require("./api/nutritionists");

app.use(express.json());

app.use("/users", usersRouter);
app.use("/nutritionists", nutritionistsRouter);

app.listen(3000, () => console.log("Holiwi 🥑"));
