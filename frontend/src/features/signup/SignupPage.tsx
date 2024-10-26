'use client'

import React from 'react'
import api from '@/api/$api'
import aspida from '@aspida/fetch'

import { Button } from '@/components/ui/button'
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { zodResolver } from '@hookform/resolvers/zod'
import { useForm } from 'react-hook-form'
import { z } from 'zod'
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/components/ui/form'
import Link from 'next/link'

const formSchema = z.object({
  email: z.string().email({ message: 'メールアドレスの形式が不適切です。' }),
  username: z.string().min(1, {
    message: 'ユーザー名は必須です。',
  }),
  password: z.string().min(6, {
    message: '6文字以上である必要があります。',
  }),
})

function SignupPage() {
  const apiClient = api(
    aspida(undefined, {
      baseURL: 'http://localhost:8080',
      credentials: 'include',
    })
  )
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      email: '',
      username: '',
      password: '',
    },
  })

  const onSubmit = async (values: z.infer<typeof formSchema>) => {
    try {
      const res = await apiClient.api.v1.signup.$post({
        body: {
          email: values.email,
          username: values.username,
          password: values.password,
        },
      })
    } catch (e) {
      console.log(e)
    }
  }

  return (
    <div
      className="bg-cover bg-center bg-no-repeat min-h-screen flex items-center justify-center"
      style={{ backgroundImage: "url('wallpaper.png')" }}
    >
      <div className="flex flex-col items-center">
        <Card className="w-[600px]">
          <div className="my-4 flex flex-col items-center ">
            <img src="tsumugi.png" width={128} height={128} alt="" />
          </div>
          <CardContent>
            <Form {...form}>
              <form
                onSubmit={form.handleSubmit(onSubmit)}
                className="space-y-8"
              >
                <FormField
                  control={form.control}
                  name="email"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>メールアドレス</FormLabel>
                      <FormControl>
                        <Input placeholder="example@email.com" {...field} />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
                <FormField
                  control={form.control}
                  name="username"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>ユーザー名</FormLabel>
                      <FormControl>
                        <Input {...field} />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
                <FormField
                  control={form.control}
                  name="password"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>パスワード</FormLabel>
                      <FormControl>
                        <Input type="password" {...field} />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
                <div className="flex justify-between items-center">
                  <Link href="/login" className="underline">
                    ログイン画面へ
                  </Link>
                  <Button type="submit">サインアップ</Button>
                </div>
              </form>
            </Form>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}

export default SignupPage
