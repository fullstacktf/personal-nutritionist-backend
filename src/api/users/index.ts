import express from "express";
import * as data from "./users.json";

const router = express.Router();
const users = data;

router.get("/", (req, res) => {
  res.json(users);
});

module.exports = router;
