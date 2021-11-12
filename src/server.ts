import express from "express";
const app = express();

const nutritionistsRouter = require("./api/nutritionists");
// import nutritionistsRouter = require("./api/nutritionists");
// require("./api/nutritionists");
// import nutritionistsRouter from "./api/nutritionists";

app.use(express.json());

app.use("/nutritionists", nutritionistsRouter);

app.listen(3000, () => console.log("Holiwi 🥑"));
