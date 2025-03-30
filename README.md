# Title
## Overview
Make a scraper which checks different sites for offers on food ingredients
Checks what can be made with the ingredients from those offers and then
suggests something to make for dinner

## Requirements
1. Be able to say yes or no to different recipes it offers

2. Define the websites for your local stores with a GUI or TUI

3. SQL Database information.
    - Add favourite flavours such as spicy or sour to influence suggestions
    - Add favourite recipes so it will choose those first if ingredients for
    them are on sale
    - Store the last X (Default 30) accepted suggestions and have the ability to let it know
    if a suggestion it made has reccently been cooked

4. Implement as frontend and backend
    - Backend. Scrapes for on-sale ingredients and comes up with different recipes to make
    from them. Storing the recipes in the database.
    - Frontend. Can access the database to get the precomputed suggestions.

5. Web Scraper
    - Finds what is on sale, and finds recipes which match at least some of the ingredients

6. Be able to grab all suggestions for the day and send them to a specific file

