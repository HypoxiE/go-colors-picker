# Go-Colors-Picker — Dynamic Wallpaper-Based Color Picker (Go)

Go-Colors-Picker is a lightweight but powerful Go tool that extracts dominant and contrasting colors from any wallpaper and generates a unified theme configuration for your Linux desktop environment.

It analyzes an image, finds frequently occurring colors, determines contrast pairs, and builds a JSON config that other tools and scripts can use to recolor your system dynamically.

Perfect for Hyprland, Eww, SwayNC, and any other program that supports theme colors.

## Features

- Reads any wallpaper image (PNG/JPG/GIF/WebP).  
- Extracts a palette of the most significant colors.  
- Determines pairs of contrasting colors for UI elements.  
- Generates a JSON config containing:  
  - `all_colors` — raw extracted colors with frequency matches  
  - `hyprland` — border and accent colors  
  - `eww` — UI, text, icon colors  
  - `swaync` — notification theme colors  
- Easy integration with scripts to:  
  - Choose a wallpaper  
  - Regenerate color config  
  - Apply theme changes across the system  

Your desktop always matches your wallpaper.

## Example Output

The generated JSON includes:  

- A list of detected colors with match counts.  
- For each supported program (Hyprland, Eww, SwayNC), a set of automatically picked main colors.  
- These colors are chosen to be both common in the wallpaper and visually distinct.  

Users can manually adjust theme colors using entries from `all_colors` if the automatic choice is not ideal.

## Usage

1. Provide a wallpaper.  
2. Run the tool with a command like: `gocp wallpaper.png`  
   or `gocp *` for all images.  
3. A config file will be created with the best matching theme colors.  
4. Use a script or your WM/compositor to apply the theme on wallpaper change.

## Why This Exists

Manually adjusting colors for Hyprland, Eww widgets, and SwayNC is tedious.  
This tool keeps your setup visually consistent and automatically adapted to your wallpaper.

## License

This project, including all past and future commits, is licensed under the **MIT License**. See the `LICENSE` file for details. You are free to use, modify, and distribute this project, including commercially, provided you include proper attribution to the original author.

---

# Go-Colors-Picker — подбор цветов по обоям (Go)

Go-Colors-Picker — лёгкая утилита на Go, которая анализирует обои, извлекает часто встречающиеся и контрастные цвета и генерирует JSON-конфиг для Hyprland, Eww, SwayNC и других программ.

Она автоматически делает вашу системную тему совпадающей с обоями.  
Если автоматический выбор цветов показался странным — можно вручную выбрать любой цвет из списка `all_colors`.

## Возможности

- Поддержка любых изображений (PNG/JPG/GIF/WebP).  
- Извлечение палитры наиболее значимых цветов.  
- Определение контрастных пар для элементов интерфейса.  
- Генерация JSON-конфига с:  
  - `all_colors` — сырые цвета с количеством вхождений  
  - `hyprland` — цвета границ и акцентов  
  - `eww` — цвета интерфейса, текста и иконок  
  - `swaync` — цвета уведомлений  
- Легкая интеграция со скриптами для:  
  - выбора обоев  
  - генерации нового конфига  
  - применения темы ко всей системе  

## Использование

1. Укажите обои.  
2. Запустите утилиту командой: `gocp wallpaper.png`  
   или `gocp *` для всех изображений.  
3. Будет создан конфиг с оптимальными цветами темы.  
4. Используйте скрипт или ваш WM/compositor для применения темы при смене обоев.

## Зачем это нужно

Ручная настройка цветов для Hyprland, Eww и SwayNC утомительна.  
Go-Colors-Picker поддерживает визуальную консистентность системы автоматически.

## Лицензия

Этот проект, включая все прошлые и будущие коммиты, распространяется под лицензией **MIT**.  
Вы можете использовать, модифицировать и распространять проект, включая коммерческое использование, при условии сохранения указания авторства.
