#!/usr/bin/python
# -*- coding: utf-8 -*-


# Author: Jie Hua 
# Date: 2023-09-22 22:29:16 
# Last Modified by:   Jie Hua 
# Last Modified time: 2023-09-22 22:29:16

from widgets.date_selector import get_date

selected_date = get_date()

if selected_date:
    print("Selected date is:", selected_date.strftime("%Y-%m-%d"))