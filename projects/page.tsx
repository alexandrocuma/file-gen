export function Page() {
	const { register, handleSubmit, formState: { errors } } = useForm<FormValues>({ resolver });
	const onSubmit = handleSubmit((data) => console.log(data));

	return (
		<form onSubmit={onSubmit}>
			<input { ...register("legal_name") } placeholder="Bill" />
			{errors?.legal_name && <p>{errors.legal_name.message}</p>}
      
			<input { ...register("name") } placeholder="Bill" />
			{errors?.name && <p>{errors.name.message}</p>}
      
			
      <input { ...register("size") } type="number"placeholder="Bill" />
			{errors?.size && <p>{errors.size.message}</p>}
		</form>
	)
}