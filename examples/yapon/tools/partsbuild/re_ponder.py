from base_ponder import BasePonder


class RoughEdgePonder(BasePonder):

    def __init__(self):

        super().__init__()

        self.dcafiles = {
            "2250": "E{}_POND_P{}",
            "1580": "E{}_POND_{}"
        }

        self.e1_signals = {}
        self.e1_signals["2250"] = {
            "e1_gap": "TN\\\\L_R1_E1GAP",
            "e1_roll_force": "TN\\\\L_R1_E1AVERLFCE",
            "e1_diff_force": "TN\\\\L_R1_E1DIFRLFCE",
        }
        self.e1_signals["1580"] = {
            "e1_gap": "TN\\\\L2_R1_E1GAP",
            "e1_roll_force": "TN\\\\L2_R1_E1AVERLFCE",
            "e1_diff_force": "TN\\\\L2_R1_E1DIFRLFCE",
        }

        self.e2_signals = {}
        self.e2_signals["2250"] = {
            "e2_gap": "TN\\\\L_R2_E2GAP",
            "e2_roll_force": "TN\\\\L_R2_E2AVERLFCE",
            "e2_diff_force": "TN\\\\L_R2_E2DIFRLFCE",
        }
        self.e2_signals["1580"] = {
            "e2_gap": "TN\\\\L2_R2_E2GAP",
            "e2_roll_force": "TN\\\\L2_R2_E2AVERLFCE",
            "e2_diff_force": "TN\\\\L2_R2_E2DIFRLFCE",
        }

    def build_df(self):

        for line in self.lines:
            for part, signal in self.e1_signals[line].items():
                for num in self.r1_nums:
                    self.df = self.df.append({
                        "LINE": line,
                        "PART": "{}{}".format(part, num),
                        "DCAFILE": self.dcafiles[line].format(1, num),
                        "SIGNAL": signal
                    }, ignore_index=True)

        for line in self.lines:
            for part, signal in self.e2_signals[line].items():
                for num in self.r2_nums:
                    self.df = self.df.append({
                        "LINE": line,
                        "PART": "{}{}".format(part, num),
                        "DCAFILE": self.dcafiles[line].format(2, num),
                        "SIGNAL": signal
                    }, ignore_index=True)









