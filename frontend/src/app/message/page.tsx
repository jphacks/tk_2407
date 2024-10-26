import { MessageCards } from '@/components/MessageCards'

const cards = [
  {
    id: 1,
    author_name: '匿名A',
    content: 'やっほー',
    color: 'bg-blue-300',
    stamps: [
      {
        type: 'thumbs_up',
        count: 4,
        is_reacted: true,
      },
      {
        type: 'heart',
        count: 2,
        is_reacted: false,
      },
    ],
  },
  {
    id: 2,
    author_name: '匿名B',
    content: 'ようこそ！',
    color: 'bg-pink-300',
    stamps: [
      {
        type: 'thumbs_up',
        count: 1,
        is_reacted: true,
      },
    ],
  },
  {
    id: 3,
    author_name: '匿名C',
    content: 'おなかすいた',
    color: 'bg-orange-300',
    stamps: [
      {
        type: 'heart',
        count: 1,
        is_reacted: false,
      },
    ],
  },
  {
    id: 4,
    author_name: '匿名D',
    content: 'また来年くるぞ〜',
    color: 'bg-blue-300',
    stamps: [],
  },
  {
    id: 5,
    author_name: '匿名E',
    content: '久々',
    color: 'bg-purple-300',
    stamps: [],
  },
  {
    id: 6,
    author_name: '匿名F',
    content: '成人しました',
    color: 'bg-yellow-300',
    stamps: [
      {
        type: 'thumbs_up',
        count: 1,
        is_reacted: false,
      },
      {
        type: 'heart',
        count: 1,
        is_reacted: false,
      },
    ],
  },
]

export default function MessagePage() {
  return <MessageCards cards={cards} />
}
