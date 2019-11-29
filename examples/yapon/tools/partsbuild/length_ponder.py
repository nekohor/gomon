from base_ponder import BasePonder


class LengthPonder(BasePonder):

    def __init__(self):

        super().__init__()

        self.tags = [
            "clg", "mfg", "flt",
            "r1", "r2",
            "fm", "f1", "f2", "f3", "f4", "f5", "f6", "f7",
            "pet", "r1dt", "r2dt", "fet", "fdt", "ct",
            "r2dw"
        ]

    def get_part_name(self, tag):
        if tag in ["f1", "f2", "f3", "f4", "f5", "f6", "f7"]:
            return "calc_length{}".format(tag[-1])
        else:
            return "{}_calc_length".format(tag)

    def build_df(self):

        for line in self.lines:
            for tag in self.tags:
                self.df = self.df.append({
                    "LINE": line,
                    "PART": self.get_part_name(tag),
                    "DCAFILE": "{}_POND".format(tag.upper()),
                    "SIGNAL": "{}\\Calc_Length".format(tag.upper())
                }, ignore_index=True)
