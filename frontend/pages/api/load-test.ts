import type { NextApiRequest, NextApiResponse } from "next";

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse
) {
  if (req.method === "POST") {
    const { target_url } = req.body;
    if (!target_url) {
      res.status(400).json({ error: "Target URL is required" });
      return;
    }

    try {
      const response = await fetch(`${process.env.BACKEND_URL}/load-test`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ target_url }),
      });

      if (!response.ok) {
        res.status(500).json({ error: "Failed to fetch from backend" });
        return;
      }

      const data = await response.json();
      res.status(200).json(data);
    } catch (error) {
      res.status(500).json({ error: "Failed to fetch from backend" });
    }
  } else {
    res.status(405).json({ error: "Method not allowed" });
  }
}
