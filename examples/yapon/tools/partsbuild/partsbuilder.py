import pandas as pd
import numpy as np
import json

from length_ponder import LengthPonder
from re_ponder import RoughEdgePonder
from rm_ponder import RouphMillPonder
from fm_ponder import FinishingMillPonder
from temperature_ponder import TemperaturePonder
from gauge_ponder import GaugePonder


class PartsBuilder():

    def __init__(self):

        self.table_name = "partTable.xlsx"

        self.ctx = {}
        self.ctx["length"] = LengthPonder()
        self.ctx["re"] = RoughEdgePonder()
        self.ctx["rm"] = RouphMillPonder()
        self.ctx["fm"] = FinishingMillPonder()
        self.ctx["temp"] = TemperaturePonder()
        self.ctx["gauge"] = GaugePonder()

    def set_fm_file_tag(self, tag="fx"):
        self.ctx["fm"].set_file_tag(tag)

    def get_fm_file_tag(self):
        return self.ctx["fm"].file_tag

    def build(self):

        df = pd.DataFrame(columns=["LINE", "PART", "DCAFILE", "SIGNAL"])

        for key, ponder in self.ctx.items():
            ponder.build_df()
            df = df.append(ponder.get_df(), ignore_index=True)

        df.to_excel(self.table_name)

    def transfer_to_json(self):

        df = pd.read_excel(self.table_name)

        part_dict = {}
        part_dict["partTable"] = []
        line_list = df["LINE"].unique()

        for i, line in enumerate(line_list):
            part_dict["partTable"].append({})
            part_dict["partTable"][i]["line"] = int(line)
            part_dict["partTable"][i]["table"] = []
            table = part_dict["partTable"][i]["table"]

            records = df.loc[df["LINE"] == line]

            for j, idx in enumerate(records.index):
                table.append({})
                table[j]["part"] = records.loc[idx, "PART"]
                table[j]["dcafile"] = records.loc[idx, "DCAFILE"]
                # print(records.loc[idx, "SIGNAL"])
                table[j]["signal"] = (
                    str(records.loc[idx, "SIGNAL"]).replace("\\\\", "\\"))

        # print(part_dict)
        with open(
            "../../Components/Tables/partTable{}.json".format(
                self.get_suffix()),
            "w",
            encoding='utf-8'
        ) as jsfile:
            json.dump(part_dict, jsfile, indent=4, ensure_ascii=False)

    def get_suffix(self):

        if self.get_fm_file_tag() == "fm":
            return "_fm"
        else:
            return ""
