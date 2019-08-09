import pandas as pd


class BasePonder():

    def __init__(self):
        self.lines = ["2250", "1580"]
        self.fm_stds = [1, 2, 3, 4, 5, 6, 7]
        self.columns = ["LINE", "PART", "DCAFILE", "SIGNAL"]
        self.df = pd.DataFrame(columns=self.columns)

        self.r1_nums = [1, 2, 3]
        self.r2_nums = [1, 2, 3, 4, 5, 6, 7, 8, 9]

        self.thermo_nums = [1, 2]

    def build_df(self):
        pass

    def get_df(self):
        return self.df

    def get_dca_file_name(self, tag):
        return "{}_POND".format(tag.upper())