const fs = require('fs');
const path = require('path');

// Функция для извлечения данных викторины из файла
function extractQuizData(filePath) {
    try {
        const content = fs.readFileSync(filePath, 'utf8');
        
        // Ищем массив quizData с разными вариантами форматирования
        const quizDataMatch = content.match(/const\s+quizData\s*=\s*(\[[\s\S]*?\]);/);
        
        if (quizDataMatch && quizDataMatch[1]) {
            try {
                // Очищаем строку, чтобы сделать её валидным JSON
                let quizDataStr = quizDataMatch[1]
                    .replace(/\/\*[\s\S]*?\*\//g, '') // Удаляем блочные комментарии
                    .replace(/\/\/.*$/gm, '') // Удаляем строчные комментарии
                    .replace(/,(\s*[}\]])/g, '$1') // Удаляем лишние запятые
                    .replace(/([\w\s]+):/g, '"$1":') // Добавляем кавычки к ключам
                    .replace(/'/g, '"'); // Заменяем одинарные кавычки на двойные

                return JSON.parse(quizDataStr);
            } catch (e) {
                console.error(`Ошибка при разборе вопросов из ${filePath}:`, e.message);
                return [];
            }
        }
        return [];
    } catch (e) {
        console.error(`Ошибка при чтении файла ${filePath}:`, e.message);
        return [];
    }
}

// Функция для обработки всех HTML файлов в директории templates
function processTemplates() {
    const templatesDir = path.join(__dirname, '../templates');
    const outputFile = path.join(templatesDir, 'general', 'index.html');
    
    // Получаем все файлы index.html, кроме того, что в папке general
    const files = [];
    
    function scanDir(dir) {
        const items = fs.readdirSync(dir, { withFileTypes: true });
        
        for (const item of items) {
            const fullPath = path.join(dir, item.name);
            
            if (item.isDirectory()) {
                scanDir(fullPath);
            } else if (item.name === 'index.html' && !fullPath.includes('general/index.html')) {
                files.push(fullPath);
            }
        }
    }
    
    scanDir(templatesDir);
    
    console.log(`Найдено файлов с вопросами: ${files.length}`);
    
    // Извлекаем и объединяем все вопросы
    const allQuestions = [];
    const processedFiles = [];
    
    files.forEach(file => {
        const questions = extractQuizData(file);
        if (questions.length > 0) {
            allQuestions.push(...questions);
            processedFiles.push({
                file: path.relative(templatesDir, file),
                count: questions.length
            });
        }
    });
    
    // Выводим статистику
    console.log('\nОбработано файлов:', processedFiles.length);
    console.log('Всего вопросов собрано:', allQuestions.length);
    console.log('\nДетали по файлам:');
    processedFiles.forEach(item => {
        console.log(`- ${item.file}: ${item.count} вопросов`);
    });
    
    try {
        // Читаем шаблон general/index.html
        let template = fs.readFileSync(outputFile, 'utf8');
        
        // Заменяем массив quizData на собранные вопросы
        const updatedContent = template.replace(
            /const\s+quizData\s*=\s*\[[\s\S]*?\];/,
            `const quizData = ${JSON.stringify(allQuestions, null, 4)};`
        );
        
        // Сохраняем обновленный файл
        fs.writeFileSync(outputFile, updatedContent, 'utf8');
        console.log(`\nФайл ${outputFile} успешно обновлен с ${allQuestions.length} вопросами`);
    } catch (e) {
        console.error('Ошибка при обновлении файла:', e.message);
    }
}

// Запускаем скрипт
processTemplates();
