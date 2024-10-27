'use client'

import React from 'react'

import { Button } from '@/components/ui/button'
import { Card, CardContent } from '@/components/ui/card'
import { Input } from '@/components/ui/input'
import { zodResolver } from '@hookform/resolvers/zod'
import { useForm } from 'react-hook-form'
import { z } from 'zod'
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/components/ui/form'
import Link from 'next/link'
import Image from 'next/image'
import apiClient from '@/apiClient'

const formSchema = z.object({
  email: z.string().email({ message: 'メールアドレスの形式が不適切です。' }),
  password: z.string().min(6, {
    message: '6文字以上である必要があります。',
  }),
})

function LoginPage() {
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      email: '',
      password: '',
    },
  })

  const onSubmit = async (values: z.infer<typeof formSchema>) => {
    try {
      const res = await apiClient.api.v1.login.$post({
        body: {
          email: values.email,
          password: values.password,
        },
      })
      console.log(res)
    } catch (e) {
      console.log(e)
    }
  }

  return (
    <div
      className="bg-cover bg-center bg-no-repeat min-h-screen w-screen flex items-center justify-center"
      style={{ backgroundImage: "url('wallpaper.png')" }}
    >
      <div className="w-full flex flex-col items-center">
        <Card className="w-[90%] md:w-[40%] sm:w-[90%]">
          <div className="my-4 flex flex-col items-center ">
            <Image
              src="/tsumugi.png"
              width={128}
              height={128}
              alt="つむぎ ロゴ"
            />
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
                  <Link href="/signup" className="underline">
                    サインアップ画面へ
                  </Link>
                  <Button type="submit">ログイン</Button>
                </div>
              </form>
            </Form>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}

export default LoginPage
