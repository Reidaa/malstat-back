#!/usr/bin/env python3

import logging
import pathlib
import sys
import time
from datetime import datetime

import httpx
import pandas as pd

logging.basicConfig(
    stream=sys.stdout,
    level=logging.INFO,
    format="%(asctime)s - %(levelname)-8s - %(message)s",
    datefmt="%Y-%m-%d %H:%M:%S",
)


class Main:
    def __init__(self) -> None:
        self.base_url = "https://api.jikan.moe/v4/"
        self.filename = "malstat.csv"
        self.mal_anime_ids = [
            52991,  # Sousou no Frieren
            5114,  # Fullmetal Alchemist: Brotherhood
            9253,  # Steins;Gate
            28977,
            38524,
            39486,
            11061,
            41467,
            9969,
            15417,
            43608,
            820,
            42938,
            34096,
            918,
            51535,
            28851,
            4181,
            35180,
            2904,
            15335,
            51009,
            37491,
            19,
            37987,
            35247,
            32281,
            40682,
            47917,
            36838,
            52198,
            45649,
            49387,
            37510,
            54492,
            40028,
            31758,
            32935,
            263,
            47778,
            199,
            48583,
            17074,
            50399,
            37521,
            1,
            50160,
            24701,
            33095,
            52034,
        ]
        self.filepath = pathlib.Path.home() / pathlib.Path(self.filename)
        self.rows = []

    def run(self):
        with httpx.Client(base_url=self.base_url) as client:
            for anime_id in self.mal_anime_ids:
                logging.info(f"Querying {anime_id}")
                r = client.get(f"anime/{anime_id}/full")
                json = r.json()
                self.rows.append(
                    {
                        "date": datetime.now(),
                        "name": json["data"]["title"],
                        "rank": json["data"]["rank"],
                        "score": json["data"]["score"],
                        "members": json["data"]["members"],
                        "favorites": json["data"]["favorites"],
                    }
                )
                time.sleep(1)
        self.to_csv()

    def to_csv(self):
        df = pd.DataFrame(self.rows)
        if not pathlib.Path(self.filepath).exists():
            df.to_csv(self.filepath, index=False)
        else:
            df.to_csv(self.filepath, mode="a", index=False, header=False)


if __name__ == "__main__":
    Main().run()
