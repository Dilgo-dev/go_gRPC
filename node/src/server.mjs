import express from "express";
import "dotenv/config";
import { createLog, getLogs } from "./grpc/logClient.mjs";

const app = express();
const PORT = process.env.PORT || 3000;

app.use(express.json());

async function log(message) {
  try {
    const response = await createLog(message);
    console.log(`Log envoyé: ${message}`);
    return response;
  } catch (error) {
    console.error(`Erreur lors de l'envoi du log: ${error.message}`);
    throw error;
  }
}

app.get("/", (req, res) => {
  log("Page d'accueil visitée");
  res.send("Hello World");
});

app.post("/log", async (req, res) => {
  try {
    const { message } = req.body;
    if (!message) {
      return res.status(400).json({ error: "Le message est requis" });
    }
    
    const response = await log(message);
    res.json(response);
  } catch (error) {
    res.status(500).json({ error: error.message });
  }
});

app.get("/logs", async (req, res) => {
  try {
    const logs = await getLogs();
    res.json(logs);
  } catch (error) {
    res.status(500).json({ error: error.message });
  }
});

app.listen(PORT, () => {
  console.log(`Server is running on port http://localhost:${PORT}`);
  
  log(`Serveur Node.js démarré sur le port ${PORT}`);
});