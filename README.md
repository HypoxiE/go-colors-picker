# README.md

# Go-Color-Picker — Dynamic Wallpaper-Based Color Picker (Go)

Go-Color-Picker is a tiny but powerful Go tool that extracts dominant and contrasting colors from any wallpaper and generates a unified theme configuration for your Linux desktop environment.

It analyzes an image, finds frequently occurring colors, determines contrast pairs, and builds a JSON config that other tools and scripts can use to recolor your system dynamically.

Perfect for Hyprland, Eww, SwayNC, and any other program that supports theme colors.

## What It Does

- Reads any wallpaper image (PNG/JPG/GIF/WebP).
- Extracts a palette of the most significant colors.
- Determines pairs of contrasting colors for UI elements.
- Generates a JSON config containing:
  - all_colors — raw extracted colors with frequency matches
  - hyprland — border and accent colors
  - eww — UI, text, icon colors
  - swaync — notification theme colors
- Makes it easy to build scripts that:
  - choose a wallpaper
  - regenerate color config
  - apply theme changes to the entire system

Your desktop always matches your wallpaper.

## Example Output (short description)

The generated JSON includes:
- A list of detected colors with match counts.
- For each supported program (Hyprland, Eww, SwayNC), a set of automatically picked "main" colors.
- These colors are chosen to be both common in the wallpaper and visually distinct from each other.

Users can manually adjust theme colors later using entries from `all_colors`.

## Usage

1. Provide a wallpaper.
2. Run the tool with a command like:
   `gocp wallpaper.png` or `gocp *` (for all images)
3. A config file will be created with the best matching theme colors.
4. Use a script or your WM compositor to apply the theme on wallpaper change.

## Why This Exists

If you use Hyprland, Eww widgets, and SwayNC together, adjusting all their colors manually is annoying.  
This tool keeps your setup visually consistent without hurting your eyes.

## License

This project, including all historical commits, is licensed under the terms of the **Creative Commons BY-NC 4.0** license.
By using this repository, you agree that all past and future versions of the code are distributed under this license.

See the LICENSE.md file for details.

---

# README (RU)

Go-Color-Picker — это утилита на Go, которая анализирует обои, извлекает часто встречающиеся и контрастные цвета и генерирует JSON-конфиг для Hyprland, Eww, SwayNC и других программ.

Она автоматически делает вашу системную тему совпадающей с обоями.  
Если автоматический выбор цветов показался странным — можно вручную выбрать любой цвет из списка `all_colors`.
