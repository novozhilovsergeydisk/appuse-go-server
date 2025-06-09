#!/usr/bin/env python3
import os
import re
import json
from pathlib import Path

def extract_questions(html_content):
    """Извлекает массив вопросов из HTML контента"""
    # Ищем массив quizData
    quiz_data_match = re.search(r'const\s+quizData\s*=\s*(\[[\s\S]*?\]);', html_content)
    if not quiz_data_match:
        return []
    
    quiz_data_str = quiz_data_match.group(1)
    
    # Извлекаем отдельные объекты вопросов
    question_matches = re.finditer(r'{\s*question\s*:\s*([\s\S]*?)\s*,\s*options\s*:\s*\[(.*?)\]\s*,\s*answer\s*:\s*(\d+)\s*}', quiz_data_str, re.MULTILINE)
    
    questions = []
    for match in question_matches:
        try:
            question = match.group(1).strip().strip('"\'')
            options_str = match.group(2)
            answer = int(match.group(3))
            
            # Извлекаем варианты ответов
            options = []
            for opt in re.findall(r'"([^"]+)"', options_str):
                options.append(opt)
            
            if len(options) > 0:
                questions.append({
                    'question': question,
                    'options': options,
                    'answer': answer
                })
        except Exception as e:
            print(f"Ошибка при разборе вопроса: {e}")
            continue
    
    return questions

def main():
    base_dir = Path(__file__).parent.parent
    templates_dir = base_dir / 'templates'
    output_file = templates_dir / 'general' / 'index.html'
    
    if not output_file.exists():
        print(f"Ошибка: Файл {output_file} не найден")
        return
    
    # Находим все index.html файлы, кроме того, что в папке general
    html_files = []
    for root, _, files in os.walk(templates_dir):
        if 'general' in root.split(os.sep):
            continue
        if 'index.html' in files:
            html_files.append(Path(root) / 'index.html')
    
    print(f"Найдено файлов с вопросами: {len(html_files)}")
    
    # Собираем все вопросы
    all_questions = []
    file_stats = []
    
    for file_path in html_files:
        try:
            with open(file_path, 'r', encoding='utf-8') as f:
                content = f.read()
            
            questions = extract_questions(content)
            if questions:
                all_questions.extend(questions)
                file_stats.append((str(file_path.relative_to(templates_dir)), len(questions)))
                print(f"  - {file_path.relative_to(templates_dir)}: {len(questions)} вопросов")
                
        except Exception as e:
            print(f"Ошибка при обработке файла {file_path}: {e}")
    
    print(f"\nВсего собрано вопросов: {len(all_questions)}")
    
    # Читаем шаблон general/index.html
    try:
        with open(output_file, 'r', encoding='utf-8') as f:
            template = f.read()
        
        # Формируем строку с массивом вопросов
        quiz_data_js = 'const quizData = ' + json.dumps(all_questions, ensure_ascii=False, indent=4) + ';'
        
        # Заменяем старый массив на новый
        updated_content = re.sub(
            r'const\s+quizData\s*=\s*\[[\s\S]*?\];',
            quiz_data_js,
            template
        )
        
        # Сохраняем обновленный файл
        with open(output_file, 'w', encoding='utf-8') as f:
            f.write(updated_content)
        
        print(f"\nФайл {output_file} успешно обновлен с {len(all_questions)} вопросами")
        
    except Exception as e:
        print(f"Ошибка при обновлении файла: {e}")

if __name__ == "__main__":
    main()
