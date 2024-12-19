export interface Interest {
  id: number;
  name: string;
  value: string;
  category: string;
}

export const interests: Interest[] = [
  // Спорт
  {
    id: 1,
    value: 'football',
    name: 'Футбол',
    category: 'Спорт'
  },
  {
    id: 2,
    value: 'basketball',
    name: 'Баскетбол',
    category: 'Спорт'
  },
  {
    id: 3,
    value: 'volleyball',
    name: 'Волейбол',
    category: 'Спорт'
  },

  // Искусство
  {
    id: 4,
    value: 'painting',
    name: 'Живопись',
    category: 'Искусство'
  },
  {
    id: 5,
    value: 'music',
    name: 'Музыка',
    category: 'Искусство'
  },
  {
    id: 6,
    value: 'photography',
    name: 'Фотография',
    category: 'Искусство'
  },

  // Технологии
  {
    id: 7,
    value: 'programming',
    name: 'Программирование',
    category: 'Технологии'
  },
  {
    id: 8,
    value: 'gaming',
    name: 'Видеоигры',
    category: 'Технологии'
  },
  {
    id: 9,
    value: 'robotics',
    name: 'Робототехника',
    category: 'Технологии'
  },

  // Наука
  {
    id: 10,
    value: 'physics',
    name: 'Физика',
    category: 'Наука'
  },
  {
    id: 11,
    value: 'chemistry',
    name: 'Химия',
    category: 'Наука'
  },
  {
    id: 12,
    value: 'biology',
    name: 'Биология',
    category: 'Наука'
  },

  // Развлечения
  {
    id: 13,
    value: 'movies',
    name: 'Кино',
    category: 'Развлечения'
  },
  {
    id: 14,
    value: 'books',
    name: 'Книги',
    category: 'Развлечения'
  },
  {
    id: 15,
    value: 'travel',
    name: 'Путешествия',
    category: 'Развлечения'
  }
];

// Получение списка всех категорий
export const getCategories = (): string[] => {
  return [...new Set(interests.map(interest => interest.category))];
};

// Получение интересов по категории
export const getInterestsByCategory = (category: string): Interest[] => {
  return interests.filter(interest => interest.category === category);
};

// Получение имени интереса по значению
export const getInterestName = (value: string): string => {
  const interest = interests.find(i => i.value === value);
  return interest ? interest.name : value;
};
