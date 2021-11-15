import express from "express";
const app = express();

const nutritionistsRouter = require("./api/nutritionists");

app.use(express.json());

app.use("/nutritionists", nutritionistsRouter);

app.listen(3000, () => console.log("Holiwi 🥑"));
