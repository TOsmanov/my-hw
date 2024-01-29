package counter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountWords(t *testing.T) {
	// Case 1
	text := ` Привет, привет мир!`
	expected := map[string]int{"мир": 1, "привет": 2}

	assert.Equal(t, expected, CountWords(text))

	// Case 2
	text = `привет мир`
	expected = map[string]int{"мир": 1, "привет": 1}

	assert.Equal(t, expected, CountWords(text))

	// Case 3
	text = `I have English on Monday, Wednesday and Thursday. English is usually the second lesson.

At the English lesson we speak, read and write. We speak about school, pupils and teachers,
about lessons, animals and nature, about our friends, sports and games.
We read books and stories about children, nature, school life and so on.
We write letters, words, sentences, exercises and dictations.

We play at English, too. We sing songs and learn poems.

I like English. I can read and write well, but I can't speak English well yet.`
	expected = map[string]int{
		"about": 4, "and": 10, "animals": 1, "at": 2, "books": 1, "but": 1,
		"can": 1, "can't": 1, "children": 1, "dictations": 1, "english": 6, "exercises": 1,
		"friends": 1, "games": 1, "have": 1, "i": 4, "is": 1, "learn": 1, "lesson": 2, "lessons": 1,
		"letters": 1, "life": 1, "like": 1, "monday": 1, "nature": 2, "on": 2, "our": 1, "play": 1,
		"poems": 1, "pupils": 1, "read": 3, "school": 2, "second": 1, "sentences": 1, "sing": 1,
		"so": 1, "songs": 1, "speak": 3, "sports": 1, "stories": 1, "teachers": 1, "the": 2,
		"thursday": 1, "too": 1, "usually": 1, "we": 6, "wednesday": 1, "well": 2,
		"words": 1, "write": 3, "yet": 1,
	}

	assert.Equal(t, expected, CountWords(text))

	// Case 4
	text = `Когда человек сознательно или интуитивно выбирает себе в жизни какую-то цель,
жизненную задачу, он невольно дает себе оценку. По тому, ради чего человек живет,
можно судить и о его самооценке - низкой или высокой.
Если человек живет, чтобы приносить людям добро, облегчать их страдания, давать людям радость,
то он оценивает себя на уровне этой своей человечности. Он ставит себе цель, достойную человека.
Только такая цель позволяет человеку прожить свою жизнь с достоинством и получить настоящую радость.
Да, радость! Подумайте: если человек ставит себе задачей увеличивать в жизни добро,
приносить людям счастье, какие неудачи могут его постигнуть? Не тому помочь?
Но много ли людей не нуждаются в помощи?
Если жить только для себя, своими мелкими заботами о собственном благополучии,
то от прожитого не останется и следа. Если же жить для других, то другие сберегут то,
чему служил, чему отдавал силы.
Можно по-разному определять цель своего существования, но цель должна быть.
Надо иметь и принципы в жизни. Одно правило в жизни должно быть у каждого человека,
в его цели жизни, в его принципах жизни, в его поведении: надо прожить жизнь с достоинством,
чтобы не стыдно было вспоминать.
Достоинство требует доброты, великодушия, умения не быть эгоистом, быть правдивым,
хорошим другом, находить радость в помощи другим.
Ради достоинства жизни надо уметь отказываться от мелких удовольствий и немалых тоже…
Уметь извиняться, признавать перед другими ошибку - лучше, чем врать.
Обманывая, человек прежде всего обманывает самого себя, ибо он думает, что успешно соврал,
а люди поняли и из деликатности промолчали.
Жизнь - прежде всего творчество, но это не значит, что каждый человек, чтобы жить,
должен родиться художником, балериной или ученым.
Можно творить просто добрую атмосферу вокруг себя. Человек может принести с собой атмосферу
подозрительности, какого-то тягостного молчания, а может внести сразу радость, свет.
Вот это и есть творчество.`

	expected = map[string]int{
		"а": 2, "атмосферу": 2, "балериной": 1, "благополучии": 1, "было": 1,
		"быть": 4, "в": 9, "великодушия": 1, "внести": 1, "вокруг": 1, "вот": 1, "врать": 1, "всего": 2,
		"вспоминать": 1, "выбирает": 1, "высокой": 1, "да": 1, "давать": 1, "дает": 1, "деликатности": 1,
		"для": 2, "добро": 2, "доброты": 1, "добрую": 1, "должен": 1, "должна": 1, "должно": 1,
		"достоинства": 1, "достоинство": 1, "достоинством": 2, "достойную": 1, "другие": 1, "другим": 1,
		"другими": 1, "других": 1, "другом": 1, "думает": 1, "его": 5, "если": 4, "есть": 1, "же": 1,
		"живет": 2, "жизненную": 1, "жизни": 7, "жизнь": 3, "жить": 3, "заботами": 1, "задачей": 1,
		"задачу": 1, "значит": 1, "и": 7, "ибо": 1, "из": 1, "извиняться": 1, "или": 3, "иметь": 1,
		"интуитивно": 1, "их": 1, "каждого": 1, "каждый": 1, "какие": 1, "какогото": 1, "какуюто": 1,
		"когда": 1, "ли": 1, "лучше": 1, "людей": 1, "люди": 1, "людям": 3, "мелкими": 1, "мелких": 1,
		"много": 1, "могут": 1, "может": 2, "можно": 3, "молчания": 1, "на": 1, "надо": 3, "настоящую": 1,
		"находить": 1, "не": 6, "невольно": 1, "немалых": 1, "неудачи": 1, "низкой": 1, "но": 3,
		"нуждаются": 1, "о": 2, "облегчать": 1, "обманывает": 1, "обманывая": 1, "одно": 1, "он": 4,
		"определять": 1, "останется": 1, "от": 2, "отдавал": 1, "отказываться": 1, "оценивает": 1, "оценку": 1,
		"ошибку": 1, "перед": 1, "по": 1, "поведении": 1, "подозрительности": 1, "подумайте": 1, "позволяет": 1,
		"получить": 1, "помочь": 1, "помощи": 2, "поняли": 1, "поразному": 1, "постигнуть": 1, "правдивым": 1,
		"правило": 1, "прежде": 2, "признавать": 1, "принести": 1, "приносить": 2, "принципах": 1, "принципы": 1,
		"прожитого": 1, "прожить": 2, "промолчали": 1, "просто": 1, "ради": 2, "радость": 5, "родиться": 1,
		"с": 3, "самого": 1, "самооценке": 1, "сберегут": 1, "свет": 1, "своего": 1, "своей": 1, "своими": 1,
		"свою": 1, "себе": 4, "себя": 4, "силы": 1, "следа": 1, "служил": 1, "собой": 1, "собственном": 1,
		"соврал": 1, "сознательно": 1, "сразу": 1, "ставит": 2, "страдания": 1, "стыдно": 1, "судить": 1,
		"существования": 1, "счастье": 1, "такая": 1, "творить": 1, "творчество": 2, "то": 4, "тоже…": 1,
		"только": 2, "тому": 2, "требует": 1, "тягостного": 1, "у": 1, "увеличивать": 1, "удовольствий": 1,
		"умения": 1, "уметь": 2, "уровне": 1, "успешно": 1, "ученым": 1, "хорошим": 1, "художником": 1,
		"цели": 1, "цель": 5, "чего": 1, "человек": 7, "человека": 2, "человеку": 1, "человечности": 1,
		"чем": 1, "чему": 2, "что": 2, "чтобы": 3, "эгоистом": 1, "это": 2, "этой": 1,
	}

	assert.Equal(t, expected, CountWords(text))
}
