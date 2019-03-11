import pandas as pd
import numpy as np
import json

# generate partTable
df = pd.read_excel("partTable.xlsx")

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
with open("partTable.json", "w", encoding='utf-8') as jsfile:
    json.dump(part_dict, jsfile, indent=4, ensure_ascii=False)


def cast_dataframe(df, col):
    cast_tag = "cast_"
    if col.startswith(cast_tag):
        new_col = col.replace(cast_tag, "")
        std_list = [1, 2, 3, 4, 5, 6, 7]
        for std in std_list:
            df[new_col + str(std)] = df[col].apply(lambda x: x + str(std))
        df.drop([col], axis=1, inplace=True)
    else:
        pass


# generate task_table
df = pd.read_excel("factorTable.xlsx")
df = df.fillna("nan")

for col in df.columns:
    cast_dataframe(df, col)

print(df)


task_dict = {}
task_dict["factorTable"] = []
tbl_list = task_dict["factorTable"]

for i, col in enumerate(df.columns):
    tbl_list.append({})
    tbl_list[i]["seriesName"] = col
    tbl_list[i]["factorList"] = []
    task_list = tbl_list[i]["factorList"]
    for task in df[col]:
        if task.startswith("nan"):
            pass
        else:
            task_list.append(task)

with open("factorTable.json", "w", encoding='utf-8') as jsfile:
    json.dump(task_dict, jsfile, indent=4, ensure_ascii=False)
