def exchange_money(budget, exchange_rate):
    return budget / exchange_rate

def get_change(budget, exchanging_value):
    return budget - exchanging_value

def get_value_of_bills(denomination, number_of_bills):
    return denomination * number_of_bills

def get_number_of_bills(budget, denomination):
    return budget // denomination


def get_leftover_of_bills(budget, denomination):
    return budget % denomination


def exchangeable_value(budget, exchange_rate, spread, denomination):
    rate = exchange_rate * (1 + spread / 100)
    exchanged = exchange_money(budget, rate)
    bills = get_number_of_bills(exchanged, denomination)
    return get_value_of_bills(denomination, bills)
