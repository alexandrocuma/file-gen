export function Page() {
	const { register, handleSubmit, formState: { errors } } = useForm<FormValues>({ resolver });
	const onSubmit = handleSubmit((data) => console.log(data));

	return (
		<form onSubmit={onSubmit}>
		</form>
	)
}