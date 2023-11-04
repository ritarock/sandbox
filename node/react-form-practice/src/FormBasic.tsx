import { useForm } from "react-hook-form"

type Data = {
  name: string,
  email: string,
  memo: string
}

export default function FormBasic() {
  const defaultValues = {
    name: "user-name",
    email: "admin@example.com",
    memo: "",
  }

  const {register, handleSubmit, formState: {errors, isDirty, isValid}} = useForm({
    defaultValues
  })

  const onsubmit = (data: Data) => console.log(data)
  const onerror = (err: unknown) => console.log(err)

  return (
    <>
      <form onSubmit={handleSubmit(onsubmit, onerror)} noValidate>
        <div>
          <label htmlFor="name">Name: </label>
          <br />
          <input id="name" type="text"
            {...register("name", {
              required: "name is required",
              maxLength: {
                value: 20,
                message: "name length <= 20"
              }
            })}
          />
          <div>{errors.name?.message}</div>
        </div>
        <div>
          <label htmlFor="email">Email: </label>
          <br />
          <input id="email" type="text"
            {...register("email", {
              required: "email is required",
              pattern: {
                value: /([a-z\d+\-.]+)@([a-z\d-]+(?:\.[a-z]+)*)/i,
                message: 'invalid email'
              }
            })}
          />
          <div>{errors.email?.message}</div>
        </div>
        <div>
          <label htmlFor="memo">Memo: </label>
          <br />
          <input id="memo" type="text"
            {...register("memo", {
              required: "memo is required",
              minLength: {
                value: 10,
                message: 'memo <= 10'
              }
            })}
          />
          <div>{errors.memo?.message}</div>
        </div>
        <div>
          <button type="submit" disabled={!isDirty || !isValid}>send</button>
        </div>
      </form>
    </>
  )
}
