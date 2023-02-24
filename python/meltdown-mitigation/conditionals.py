"""Functions to prevent a nuclear meltdown."""


def is_criticality_balanced(temperature, neutrons_emitted):
    return temperature < 800 and neutrons_emitted > 500 and (neutrons_emitted * temperature) < 500000

def reactor_efficiency(voltage, current, theoretical_max_power):

    power = voltage * current
    efficiency = (power / theoretical_max_power) * 100

    if efficiency >= 80:
        return 'green'
    if efficiency >= 60:
        return 'orange'
    if efficiency >= 30:
        return 'red'

    return 'black'

def fail_safe(temperature, neutrons_produced_per_second, threshold):

    activity = temperature * neutrons_produced_per_second

    if activity < threshold * 0.9:
        return 'LOW'
    if activity > threshold * 1.1:
        return 'DANGER'

    return 'NORMAL'
