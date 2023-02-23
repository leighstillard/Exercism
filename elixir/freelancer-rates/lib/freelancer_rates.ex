

defmodule FreelancerRates do
  def daily_rate(hourly_rate) do
    hourly_rate * 8.0
  end

  def apply_discount(before_discount, discount) do
    before_discount * (1 - discount / 100)
  end

  def monthly_rate(hourly_rate, discount) do
    trunc(Float.ceil(apply_discount(22 * 8.0 * hourly_rate, discount),0))
  end

  @spec days_in_budget(number, number, number) :: float
  def days_in_budget(budget, hourly_rate, discount) do
    dayrate = apply_discount(daily_rate(hourly_rate), discount)
    Float.floor(budget / dayrate, 1)
  end
end
